function goToUrl(url) {
  window.location = `${window.location.origin}${url}`;
}

function showMessage(text = '', type = 'success') {
  let actionTextColor;

  switch (type) {
    case 'error':
      actionTextColor = '#f44336';
      break;
    case 'info':
      actionTextColor = '#f44336';
      break;
    default:
      actionTextColor = '#4caf50';
  }

  Snackbar.show({
    pos:'bottom-left',
    actionText: 'Fechar',
    actionTextColor,
    text
  });
}
