window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
      {"url":"grill/v1/grill_service.swagger.json","name":"grill_service.swagger.json"},
      {"url":"user/v1/user_service.swagger.json","name":"user_service.swagger.json"},
      {"url": "agent/v1/agent_service.swagger.json", "name": "agent_service.swagger.json"}
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
