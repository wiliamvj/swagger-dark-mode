package custom

import "fmt"

var CustomJS = fmt.Sprintf(`
   // set custom title
    document.title = 'Swagger Dark Mode With Go';

    // set custom favicon
    const link = document.createElement('link');
    link.rel = 'icon';
    link.type = 'image/x-icon';
    link.href = 'data:image/png;base64,%s';
    document.head.appendChild(link);

    // set custom logo
    const image = document.querySelector('.link img');
    const base64URL = 'data:image/png;base64,%s';
    image.src = base64URL;

    // dark mode
    const style = document.createElement('style');
    style.innerHTML = %s;
    document.head.appendChild(style);
  `, CustomLogo, CustomLogo, "`"+customCSS+"`")
