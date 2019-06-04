import 'dart:html';

import '../formstate.dart';

class ContactItem extends FormState {
  EmailInputElement _email;
  TelephoneInputElement _phone;

  ContactItem(
      String formID, String emailElem, String phoneElem, String submitID)
      : super(formID, submitID) {
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
  }

  String get email {
    return _email.value;
  }

  String get phone {
    return _phone.value;
  }

  Object toJson() {
    return {"Email": email, "Phone": phone};
  }
}