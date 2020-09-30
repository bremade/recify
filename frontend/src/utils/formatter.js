function formatTime(minutes) {
    if (minutes < 60) {
      return `${minutes} min`;
    } else {
      const h = Math.floor(minutes / 60);
      const m = minutes % 60;

      if (m > 0) {
        return `${h} h ${m} min`;
      } else {
        return `${h} h`;
      }
    }
  }

  function formatPrice(price) {
    const eur = Math.floor(price);
    const cent = (price - eur) * 100;

    return `${eur}.${String(cent).padStart(2, '0')}`;
  }

export default {
    formatTime,
    formatPrice
};
