import React from "react";
import { createRoot } from "react-dom/client";
import htm from "htm";

import { BrowserRouter } from "react-router-dom";
import { App } from "./App.js";

const html = htm.bind(React.createElement);

const root = createRoot(document.getElementById("root"));

root.render(html`<${BrowserRouter}><${App} /></${BrowserRouter}>`);
