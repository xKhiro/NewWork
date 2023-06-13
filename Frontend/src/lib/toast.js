export function showToast(message) {
  const toastBody = document.getElementById('toast');
  const toastMessage = document.getElementById('toastBody');
  toastMessage.innerHTML = message;

  const toast = new bootstrap.Toast(toastBody);
  toast.show();
}
