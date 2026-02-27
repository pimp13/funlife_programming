use futures::future::join_all;
use reqwest;
use tokio::{fs::File, io::AsyncWriteExt};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // let filename = "rust-logo.png";
    // let url = "https://www.rust-lang.org/logos/rust-logo-512x512.png";
    // download(url, filename).await?;
    // Ok(())

    let urls = vec![
        ("http://picsum.photos/1920/1080", "img1.png"),
        ("http://picsum.photos/1920/1080", "img2.jpg"),
        ("http://picsum.photos/1920/1080", "img3.jpg"),
        ("http://picsum.photos/1920/1080", "img4.jpg"),
        ("http://picsum.photos/1920/1080", "img5.jpg"),
        ("http://picsum.photos/1920/1080", "img6.jpg"),
    ];

    // for u in urls {
    //     download(u.0, u.1).await?;
    // }

    let fut = urls.iter().map(|(url, name)| download(url, name));
    let results = join_all(fut).await;
    for result in results {
        if let Err(e) = result {
            eprintln!("failed to download async file from urls: {}", e)
        }
    }

    Ok(())
}

async fn download(url: &str, filename: &str) -> Result<(), Box<dyn std::error::Error>> {
    let res = reqwest::get(url).await?;
    let bytes = res.bytes().await?;

    // use tokio file
    let mut file = File::create(filename).await?;
    file.write_all(&bytes).await?;

    // use std file
    // let mut file = File::create(filename)?;
    // file.write_all(&bytes)?;

    println!("Downloded {} is successfully!", filename);
    Ok(())
}
