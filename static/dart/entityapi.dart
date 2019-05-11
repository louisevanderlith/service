import 'dart:convert';
import 'dart:html';

import 'pathlookup.dart';

Future<String> createEntity(Object data) async {
  var url = await buildPath("Entity.API", "info", new List<String>());
  return HttpRequest.requestCrossOrigin(url,
      method: "POST", sendData: jsonEncode(data));
}