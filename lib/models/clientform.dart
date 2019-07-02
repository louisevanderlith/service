import 'dart:convert';
import 'dart:html';

import '../entityapi.dart';
import '../formstate.dart';
import 'contactitem.dart';

class ClientForm extends FormState {
  TextInputElement _name;
  ContactItem _contact;
  TextInputElement _identification;
  ParagraphElement _error;

  ClientForm(String formID, String nameElem, String emailElem, String phoneElem, String identityElem,
      String submitID)
      : super(formID, submitID) {
    _name = querySelector(nameElem);
    _contact = new ContactItem(formID, emailElem, phoneElem, submitID);
    _identification = querySelector(identityElem);

    querySelector(submitID).onClick.listen(onSend);
  }

  String get name {
    return _name.value;
  }

  ContactItem get contact {
    return _contact;
  }

  String get identification {
    return _identification.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var result = await createEntity(name, contact, identification);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        window.alert(obj['Data']);
        disableSubmit(false);
      } else {
        _error.text = obj['Error'];
      }
    }
  }
}
