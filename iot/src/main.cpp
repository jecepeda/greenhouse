#include <WiFi.h>
#include "time.h"
#include "Credentials.hpp"
#include <HTTPClient.h>
#include <ArduinoJson.h>

void setupWifi();
void loginDevice();
void refreshTokens();
void sendTelemetry();

// auth values
String accessToken = "";
String refreshToken = "";

void setup()
{
  Serial.begin(115200);
  setupWifi();
  loginDevice();
}

void loop()
{
  delay(1000);
  // refreshTokens();
  sendTelemetry();
}

void sendTelemetry()
{
  refreshTokens();
  Serial.println("setting up data");
  DynamicJsonDocument doc(1024);
  doc["temperature"] = 12.345;
  doc["humidity"] = 13.45;
  doc["heater_enabled"] = true;
  doc["humidifier_enabled"] = true;
  Serial.println("finish setting up data");
  if (WiFi.status() == WL_CONNECTED)
  {
    Serial.println("setting up http client");
    HTTPClient http;
    http.useHTTP10(true);
    http.begin(getHost() + "/v1/device/" + String(getDeviceID()) + "/monitoring/data");
    http.addHeader("Authorization", "Bearer " + accessToken);
    Serial.println("finish setting up http client");
    String body;
    serializeJson(doc, body);
    Serial.println("finish serialicing");
    int responseCode = http.POST(body);
    Serial.println("finish request");
    Serial.println(responseCode);
    http.end();
  }
  else
  {
    setupWifi();
  }
}

void refreshTokens()
{
  Serial.println("refresh tokens");
  if (WiFi.status() == WL_CONNECTED)
  {
    HTTPClient http;
    http.useHTTP10(true);
    http.begin(getHost() + "/v1/device/refresh");
    http.addHeader("Authorization", "Bearer " + refreshToken, false, false);
    int responseCode = http.POST("");
    if (responseCode == 200)
    {
      DynamicJsonDocument response(2048);
      deserializeJson(response, http.getStream());
      accessToken = response["access_token"].as<String>();
      refreshToken = response["refresh_token"].as<String>();
      Serial.println("Refresh token has been obtained");
    }
    http.end();
  }
}

void loginDevice()
{
  HTTPClient http;
  http.useHTTP10(true);
  if (WiFi.status() == WL_CONNECTED)
  {
    DynamicJsonDocument doc(1024);
    doc["device"] = getDeviceID();
    doc["password"] = getPassword();

    String json;
    serializeJson(doc, json);

    http.begin(getHost() + "/v1/device/login");
    int responseCode = http.POST(json);
    if (responseCode == 200)
    {
      DynamicJsonDocument response(2048);
      deserializeJson(response, http.getStream());
      accessToken = response["access_token"].as<String>();
      refreshToken = response["refresh_token"].as<String>();
    }
  }
  else
  {
    setupWifi();
  }
  http.end();
}

void setupWifi()
{
  Serial.print("Connecting to ");
  Serial.println(getSSID());
  WiFi.begin(getSSID(), getSSIDPassword());
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("WiFi connected.");
  Serial.print("Connected to WiFi network with IP Address: ");
  Serial.println(WiFi.localIP());
}