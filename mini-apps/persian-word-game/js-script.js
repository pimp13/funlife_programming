#!/bin/env node

const fs = require('fs');

fs.readFile('./persian_dict_19k.csv', 'utf8', (err, file) => {
  if (err) {
    console.error(err);
    return;
  }

  

  const word = file.split(":")[0];
  console.log(word);

});