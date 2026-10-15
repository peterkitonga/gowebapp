const toggleToast = (message, type) => {
  let toastElClasses = 'toast align-items-center border-0';
  let htmlMessage = '';

  if (type === 'success') {
    toastElClasses += ' text-success-emphasis bg-success-subtle';
    htmlMessage = `<i class="bi bi-check-circle-fill"></i> <span>${message}</span>`;
  } else if (type === 'error') {
    toastElClasses += ' text-danger-emphasis bg-danger-subtle';
    htmlMessage = `<i class="bi bi-x-octagon-fill"></i> <span>${message}</span>`;
  } else if (type === 'warning') {
    toastElClasses += ' text-warning-emphasis bg-warning-subtle';
    htmlMessage = `<i class="bi bi-exclamation-triangle-fill"></i> <span>${message}</span>`;
  } else if (type === 'info') {
    toastElClasses += ' text-info-emphasis bg-info-subtle';
    htmlMessage = `<i class="bi bi-info-circle-fill"></i> <span>${message}</span>`;
  } else {
    toastElClasses += ' text-dark-emphasis bg-dark-subtle';
    htmlMessage = `<i class="bi bi-question-circle-fill"></i> <span>${message}</span>`;
  }

  const toastOptions = {
    delay: 10000,
  };
  const toastEl = document.getElementById('toast-el');
  const toastElBody = document.getElementById('toast-el-body');
  const toastInstance = new bootstrap.Toast(toastEl, toastOptions);

  toastEl.className = toastElClasses;
  toastElBody.innerHTML = htmlMessage;
  toastInstance.show();
};
