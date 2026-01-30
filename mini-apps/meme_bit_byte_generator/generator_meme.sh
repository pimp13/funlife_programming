#!/bin/bash

word="${1:-rabbit}"

# 1. دریافت URL تصویر
echo "Getting image URL for: $word"
img_url=$(go run ./mime.go "$word")

# بررسی موفقیت‌آمیز بودن
if [ $? -ne 0 ] || [ -z "$img_url" ]; then
    echo "Error: Failed to get image URL"
    exit 1
fi

echo "Image URL: $img_url"

# 2. دانلود تصویر
output_img="temp.png"
wget "$img_url" -O "$output_img" 2>/dev/null

# بررسی دانلود
if [ $? -ne 0 ] || [ ! -f "$output_img" ]; then
    echo "Error: Failed to download image"
    exit 1
fi

echo "Image downloaded successfully"

# 3. تغییر سایز تصویر
convert "$output_img" -resize x200 "$output_img"

# 4. تبدیل bit به byte
wordbyte=$(echo "$word" | sed 's/bit$/byte/')
echo "Word: $word, Byte version: $wordbyte"

# font -font Arial 
# 5. ایجاد top.png (یک تصویر با عنوان)
convert "$output_img" \
    -gravity north -background white -splice 0x50 \
    -pointsize 24 -fill black \
    -annotate +0+20 "$word" \
    "top.png"

# 6. ایجاد bottom.png (8 تصویر تایل شده با عنوان)
# اول 8 کپی از تصویر ایجاد می‌کنیم
montage "$output_img" "$output_img" "$output_img" "$output_img" \
        "$output_img" "$output_img" "$output_img" "$output_img" \
        -tile 4x2 -geometry +5+5 \
        "temp_grid.png"

# سپس عنوان اضافه می‌کنیم
convert "temp_grid.png" \
    -gravity north -background white -splice 0x50 \
    -font Arial -pointsize 24 -fill black \
    -annotate +0+20 "$wordbyte" \
    "bottom.png"

# 7. ترکیب top و bottom
convert "top.png" "bottom.png" -append "${word}.png"

# 8. تمیزکاری
rm -f "temp_grid.png" "top.png" "bottom.png" "temp.png"

echo "Meme created successfully: ${word}.png"