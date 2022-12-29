import { createPostStatusCmd } from "./create-post-status-cmd.js";
import { statusUpdateService } from "./status-update-service.js";
import { Invoker } from "./invoker.js";

const invoker = new Invoker();

const command = createPostStatusCmd(statusUpdateService, 'HI!');

invoker.run(command);
invoker.undo(command);
invoker.delay(command, 1000 * 3);
// invoker.runRemotely(command);
