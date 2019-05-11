import 'dart:html';

class FormState {
  FormElement _form;
  ButtonElement _sendBtn;

  FormState(String formID, String submitID) {
    _form = querySelector(formID);
    _sendBtn = querySelector(submitID);

    disableSubmit(true);

    _form.onKeyUp.listen(pressEnter);
  }

  FormElement get form {
    return _form;
  }

  ButtonElement get submit {
    return _sendBtn;
  }
  
  bool isFormValid() {
    return _form.checkValidity();
  }

  void disableSubmit(bool disable) {
    _sendBtn.disabled = disable;
  }

  void registerFormElements(List<Element> elements) {
    for (var elem in elements) {
      var runtime = elem.runtimeType.toString();
      Function validFunc;
      if (runtime == "InputElement") {
        validFunc = validateElement;
      }

      if (runtime == "TextAreaElement") {
        validFunc = validateAreaElement;
      }

      if (validFunc != null) {
        elem.onBlur.listen((e) => validFunc(e, elem));
      }
    }
  }

  void validateElement(Event e, InputElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    disableSubmit(!isFormValid());
  }

  void validateAreaElement(Event e, TextAreaElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    disableSubmit(!isFormValid());
  }

  void pressEnter(KeyboardEvent e) {
    if (e.key != 'Enter') {
      return;
    }

    e.preventDefault();
    _sendBtn.click();
  }
}
