import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:Service.APP/models/contactitem.dart';

import 'pathlookup.dart';

Future<HttpRequest> createEntity(String name, ContactItem contact, String identification) async {
  final url = await buildPath("Entity.API", "info", new List<String>());
  
  final data = jsonEncode({
      "Name": name,
      "Contact": contact,
      "Identification": identification
    });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader("Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}