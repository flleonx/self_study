import fs from 'fs';
import path from 'path';
import superagent from 'superagent';
import mkdirp from 'mkdirp';
import { urlToFilename, getPageLinks } from './utils.js';

function saveFile(filename, contents, cb) {
  mkdirp(path.dirname(filename), err => {
    if (err) {
      return cb(err);
    }

    fs.writeFile(filename, contents, cb)
  })
};

function download(url, filename, cb) {
  console.log(`Downloading ${url}`);
  superagent.get(url).end((err, res) => {
    if (err) {
      return cb(err);
    }

    saveFile(filename, res.text, err => {
      if (err) {
        return cb(err);
      }
      console.log(`Downloaded and saved: ${url}`);
      cb(null, res.text);
    })
  })
};

function spiderLinks(currentUrl, body, nesting, queue) {
  console.log('spiderLinks starts');
  if (nesting === 0) {
    console.log('Nesting 0 condition');
    return;
  };

  const links = getPageLinks(currentUrl, body);
  console.log('LINKS', links);
  if (links.length === 0) {
    console.log('No links condition');
    return;
  };

  links.forEach(link => spider(link, nesting - 1, queue));
}

export function spiderTask(url, nesting, queue, cb) {
  const filename = urlToFilename(url);
  console.log('Filename: ', filename);
  console.log('Starting spider');
  fs.readFile(filename, 'utf8', (err, fileContent) => {
    if (err) {
      if (err.code !== 'ENOENT') {
        return cb(err);
      };

      // The file doesn't exist, so let's download it
      return download(url, filename, (err, requestContent) => {
        if (err) {
          return cb(err);
        };
        spiderLinks(url, requestContent, nesting, queue);
        console.log('Return spiderlinks');
        return cb();
      });
    };

    // The file already exists, let's process the links
    console.log('File exists spiderlinks');
    spiderLinks(url, fileContent, nesting, queue);
    console.log('Return file exists spiderlinks');
    return cb();
  });
};

const spidering = new Set();
export function spider(url, nesting, queue) {
  // Avoid race conditions
  if (spidering.has(url)) {
    return;
  }
  spidering.add(url);
  queue.pushTask((done) => {
    spiderTask(url, nesting, queue, done);
  });
};
