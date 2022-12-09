import { Config } from "./config.js";
import { jsonStrategy, iniStrategy } from "./strategies.js";

/*
  The structure of the Strategy pattern may look similar to that of
  the Adapter pattern. However, there is a substantial difference between
  the two. The adapter object does not add any behavior to the adaptee;
  it just makes it avaliable under another interface. This can also require
  some extra logic to be implemented to convert one interface into another,
  but this logic is limited to this task only. In the Strategy pattern,
  however, the context and the strategy implement some kind of logic and
  both are essential to build the final algorithm (when combined together).
*/

async function main() {
  const iniConfig = new Config(iniStrategy);
  await iniConfig.load('samples/conf.ini');
  iniConfig.set('book.nodejs', 'design patterns');
  await iniConfig.save('samples/conf_mod.ini');

  const jsonConfig = new Config(jsonStrategy);
  await jsonConfig.load('samples/conf.json');
  jsonConfig.set('book.nodejs', 'design patterns');
  await jsonConfig.save('samples/conf_mod.json');
}

main();
