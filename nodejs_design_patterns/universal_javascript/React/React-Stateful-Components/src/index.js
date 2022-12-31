import React from "react";
import { createRoot } from "react-dom/client";
import htm from "htm";
import { App } from "./App.js";
import { ErrorBoundary } from "./ErrorBoundary.js";

const html = htm.bind(React.createElement);

const root = createRoot(document.getElementById("root"));

root.render(html`<${ErrorBoundary}><${App} /></${ErrorBoundary}>`);
