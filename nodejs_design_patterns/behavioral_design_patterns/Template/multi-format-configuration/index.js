import { JsonConfig } from "./json-config.js";
import { IniConfig } from "./ini-config.js";

async function main() {
  const jsonConfig = new JsonConfig();
  await jsonConfig.load("samples/conf.json");
  jsonConfig.set("nodejs", "design patterns");
  await jsonConfig.save("samples/conf_mod.json");

  const iniConfig = new IniConfig();
  await iniConfig.load("samples/conf.ini");
  iniConfig.set("nodejs", "design patterns");
  await iniConfig.save("samples/conf_mod.ini");
}

main();
