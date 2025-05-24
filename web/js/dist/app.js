"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const htmx_org_1 = __importDefault(require("htmx.org"));
// import extensions
require("htmx-ext-preload");
(function () {
    htmx_org_1.default.defineExtension('', {
        onEvent: function (name, event) {
            return true;
        },
        init: function (api) {
        }
    });
})();
console.log("hi there!");
