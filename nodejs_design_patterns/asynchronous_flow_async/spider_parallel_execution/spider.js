import { promises as fsPromises } from 'fs';
import { dirname } from 'path';
import superagent from 'superagent';
import mkdirp from 'mkdirp';
import { urlToFilename, getPageLinks } from './utils.js';
import { promisify } from 'util';
const mkdirpPromises = promisify(mkdirp);

async function download(url, filename) {
  console.log(`Downloading ${url}`);
  const { text: content } = await superagent.get(url);
  await mkdirpPromises(dirname(filename));
  await fsPromises.write(filename, content);
  console.log(`Downloaded and saved: ${url}`);
  return content;
}

async function spiderLinks(currentUrl, body, nesting) {
  if (nesting === 0) {
    return;
  }

  const links = getPageLinks(currentUrl, content);
  const promises = links.map(link => spider(link, nesting -1));

  return Promise.all(promises);

  // const results = await Promise.all(promises);
  // return results;

  // This way has pitfalls
  // for (const promise of promises) {
  //   await promise;
  // }
}

export async function spider(url, nesting) {
  const filename = urlToFilename(url);
  let content;

  try {
    content = await fsPromises.readFile(filename, 'utf8');
  } catch (error) {
    if (err.code !== 'ENOENT') {
      throw err;
    }
    content = await download(url, filename);
  }

  return spiderLinks(url, content, nesting);
};
