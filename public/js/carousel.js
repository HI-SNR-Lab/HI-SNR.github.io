function isVisible(el) {
    while (el) {
        if (el === document) {
            return true;
        }

        const style = window.getComputedStyle(el);

        if (style.display === 'none') return false;
        if (style.visibility === 'hidden') return false;
        if (parseFloat(style.opacity) === 0) return false;

        el = el.parentNode;
    }

    return false;
}

setInterval(function () {
    let j = 0;

    const elements = document.querySelectorAll(
        '.carousel .carousel__control--forward'
    );

    for (let i = elements.length - 1; i >= 0; i--) {
        if (isVisible(elements[i])) {
            j = i;
            break;
        }
    }

    if (elements[j]) {
        elements[j].click();
    }
}, 5000);