import { UrlBuilder } from './url-builder.js';

/*
  Using a builder that is separate from the target class has the
  advantage of always producing instances that are guaranteed to be
  in a consistent state
*/

const url = new UrlBuilder()
  .setProtocol('https')
  .setAuthentication('user', 'pass')
  .setHostname('example.com')
  .build()

console.log(url.toString());
