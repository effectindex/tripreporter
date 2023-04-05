// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

const { defineConfig } = require("@vue/cli-service")
const path = require("path")

require("dotenv").config({
  path: path.join(__dirname, "..", ".env")
});

module.exports = defineConfig({
  transpileDependencies: true,
  assetsDir: "static",
  devServer: {
    port: process.env.DEV_PORT,
  }
})
