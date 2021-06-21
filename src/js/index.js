import React, { StrictMode } from "react";
import ReactDOM from "react-dom";

import App from "./app";

const rootElement = document.getElementById("app");
ReactDOM.render(
  <StrictMode>
    <App />
  </StrictMode>,
  rootElement
);