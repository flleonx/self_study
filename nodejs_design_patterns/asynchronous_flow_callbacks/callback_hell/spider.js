import fs from 'fs';
import path from 'path';
import superagent from 'superagent';
import mkdirp from 'mkdirp';
import { urlToFilename } from './utils.js';

export function spider(url, callback) {
  const filename = urlToFilename(url);
  fs.access(filename, err => {
    if (err && err.code === 'ENOENT') {
      console.log(`Downloading ${url} into ${filename}`);
      superagent.get(url).end((err, res) => {
        if (err) {
          callback(err);
        } else {
          mkdirp(path.dirname(filename), err => {
            if (err) {
              callback(err);
            } else {
              fs.writeFile(filename, res.text, err => {
                if (err) {
                  callback(err);
                } else {
                  callback(null, filename, true);
                }
              })
            }
          })
        }
      })
    } else {
      callback(null, filename, false);
    }
  })
}
