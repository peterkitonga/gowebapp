const CURRENT_DATE = new Date();
const DATEPICKER_OPTIONS = {
  allowInputToggle: true,
  useCurrent: true,
  viewDate: CURRENT_DATE,
  display: {
    calendarWeeks: true,
    keepOpen: false,
    buttons: {
      today: true,
      clear: true,
      close: true,
    },
    icons: {
      type: 'icons',
      time: 'bi bi-clock',
      date: 'fbi bi-calendar-event',
      up: 'bi bi-arrow-up-circle',
      down: 'bi bi-arrow-down-circle',
      previous: 'bi bi-chevron-double-left',
      next: 'bi bi-chevron-double-right',
      today: 'bi bi-calendar-check',
      clear: 'bi bi-trash',
      close: 'bi bi-x-circle',
    },
    components: {
      clock: false,
    },
  },
  localization: {
    format: 'L',
    dateFormats: {
      LTS: 'h:mm:ss T',
      LT: 'h:mm T',
      L: 'yyyy-MM-dd',
      LL: 'MMMM d, yyyy',
      LLL: 'yyyy-MM-dd h:mm T',
      LLLL: 'dddd, MMMM d, yyyy h:mm T',
    },
  },
  restrictions: {
    minDate: CURRENT_DATE,
  },
};
