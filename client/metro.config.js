/* eslint-env node */

const { getDefaultConfig } = require("expo/metro-config");
// const { mergeConfig } = require("@react-native/metro-config");
// const { mergeConfig } = require("metro-config");
const { mergeConfig } = require("@react-native/metro-config");
const { withNativeWind } = require("nativewind/metro");

const defaultConfig = getDefaultConfig(__dirname);

const config = {
  resolver: {
    // Connect-ES and Protobuf-ES use package exports
    // (https://nodejs.org/docs/latest-v12.x/api/packages.html#packages_exports).
    //
    // We need to enable support for them in Metro. See https://metrobundler.dev/docs/package-exports/
    unstable_enablePackageExports: true,
    // unstable_enableSymlinks: true,
  },
  transformer: {
    unstable_allowRequireContext: true,
  },
};

const mergedConfig = mergeConfig(defaultConfig, config);

module.exports = withNativeWind(mergedConfig, { input: "./global.css" });
