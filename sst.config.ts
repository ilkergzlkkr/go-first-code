/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "go-first-code",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
    };
  },
  async run() {
    const api = new sst.aws.ApiGatewayV2("GoApi", {
      domain:
        $app.stage === "production"
          ? {
              name: "go.illker.dev",
              dns: sst.cloudflare.dns(),
            }
          : undefined,
    });
    api.route("$default", {
      handler: ".",
      runtime: "go",
    });

    return {
      api: api.url,
    };
  },
});
