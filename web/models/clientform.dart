import 'dart:async';
import 'dart:html';

import '../entityapi.dart';
import '../formstate.dart';
import 'contactitem.dart';

class ClientForm extends FormState {
  TextInputElement _name;
  ContactItem _contact;

  ClientForm(String formID, String nameElem, String emailElem, String phoneElem,
      String submitID)
      : super(formID, submitID) {
    _name = querySelector(nameElem);
    _contact = new ContactItem(formID, emailElem, phoneElem, submitID);

    querySelector(submitID).onClick.listen(onSend);
    registerFormElements([_name]);
  }

  String get name {
    return _name.value;
  }

  ContactItem get contact {
    return _contact;
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend().then((obj) {
        disableSubmit(false);
      });
    }
  }

  Future submitSend() async {
    var obj = {
      "Name": name,
      "Contact": contact,
    };

    return await createEntity(obj);
  }
}
