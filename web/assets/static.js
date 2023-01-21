
function redirecionar(title, id, url){
    var mensagem = new SpeechSynthesisUtterance();
    mensagem.text = title
    speechSynthesis.speak(mensagem);
    mensagem.onend = () => { location.href = url+"?id="+ id}
}

function redirecionarLink(obj, url){
  var mensagem = new SpeechSynthesisUtterance();
  mensagem.text = "Abrindo "+obj.text
  speechSynthesis.speak(mensagem);
  mensagem.onend = () => { location.href = url+"?id="+obj.title }
}

function lerTexto(title){
  var mensagem = new SpeechSynthesisUtterance();
  mensagem.text = title
  speechSynthesis.speak(mensagem);
}