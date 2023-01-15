<template>
  <div class="LayoutDefault">
    <nav class="LayoutDefault__nav">
      <router-link class="LayoutDefault__item" to="/">Home</router-link>
      <router-link class="LayoutDefault__item" to="/about">About</router-link>
      <router-link class="LayoutDefault__item" to="/signup">Signup</router-link>
      <router-link class="LayoutDefault__item" to="/login">Login</router-link>
    </nav>
    <main class="LayoutDefault__main">
      <slot/>
    </main>
    <footer class="LayoutDefault__footer">
      <div class="LayoutDefault__footer_box">
        <div class="LayoutDefault__footer_social">
          <a href="https://www.youtube.com/c/Josikinz">
            <img src="../assets/svg/sm-icon-youtube.svg" alt="YouTube" width="32" height="32">
          </a>
          <a href="/discord">
            <img src="../assets/svg/sm-icon-discord.svg" alt="Facebook" width="32" height="32">
          </a>
          <a href="https://reddit.com/r/replications">
            <img src="../assets/svg/sm-icon-reddit.svg" alt="Reddit" width="32" height="32">
          </a>
          <a href="https://github.com/effectindex/tripreporter">
            <img src="../assets/svg/sm-icon-github.svg" alt="GitHub" width="32" height="32">
          </a>
        </div>
      </div>
      <div class="LayoutDefault__footer_box">
        <div class="LayoutDefault__footer_text">
          &copy; 2022, Subjective Effect Documentation, 5HT2 and contributors.
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "LayoutDefault"
}
</script>

<style>
@import url(@/assets/css/default.css);
</style>

<style scoped>
*, ::after, ::before {
  box-sizing: border-box;
  margin: 0;
}

.LayoutDefault {
  font-family: "Titillium Web", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

.LayoutDefault__nav {
  font-family: "Proxima Nova", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  background-color: var(--tr-default-primary-background);
  height: 87px;
  padding: 20px;
  display: flex;
  flex-grow: 1;
  justify-content: flex-end;
}

.LayoutDefault__nav a {
  color: var(--tr-default-alt-text);
  text-decoration: none;
  text-transform: uppercase;
  font-weight: 300;
  letter-spacing: 2px;
  transition: color .15s ease-in-out;
  margin-left: 1em;
  padding: 10px 0;
  font-size: 13px;
  white-space: pre;
}

.LayoutDefault__nav a.router-link-exact-active {
  color: var(--tr-default-primary);
}

.LayoutDefault__nav > a:hover {
  color: var(--tr-default-primary-text);
  text-shadow: var(--tr-default-primary-shadow);
}

/* fix positioning and alignment of nav items */
.LayoutDefault__item {
  position: relative;
  line-height: 33px;
  align-self: center;
  transition: color .5s ease;
  margin: 0 1em;
}

/* pin footer to bottom */
.LayoutDefault__footer {
  color: var(--tr-default-alt-darker-text);
  background-color: var(--tr-default-primary-background);
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 90px;
}

/* vertically center footer text */
.LayoutDefault__footer_text {
  position: relative;
  top: 50%;
  transform: translateY(-50%);
}

/* move social icons further down */
.LayoutDefault__footer_social {
  position: relative;
  top: 20%;
}

/* prettify social icons + animate hover */
.LayoutDefault__footer_social > a {
  text-decoration: none;
  display: inline-block;
  margin: .3em;
  transition: all .2s ease-in-out;
  height: 32px;
  width: 32px;
}

/* expand social icons on hover */
.LayoutDefault__footer_social > a:hover {
  transform: scale(1.2);
}

/* fix size of box */
.LayoutDefault__footer_box {
  height: 50%;
}

/* TODO: Nicer way to do this */
/* latin-ext */
@font-face {
  font-family: 'Titillium Web';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPAcZTIAOhVxoMyOr9n_E7fdMbWAaxWXr0.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Titillium Web';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPAcZTIAOhVxoMyOr9n_E7fdMbWD6xW.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* latin-ext */
@font-face {
  font-family: 'Titillium Web';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPFcZTIAOhVxoMyOr9n_E7fdMbetIlDb5yciWM.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Titillium Web';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPFcZTIAOhVxoMyOr9n_E7fdMbetIlDYZyc.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* latin-ext */
@font-face {
  font-family: 'Titillium Web';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPecZTIAOhVxoMyOr9n_E7fdM3mDbRS.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Titillium Web';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPecZTIAOhVxoMyOr9n_E7fdMPmDQ.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* latin-ext */
@font-face {
  font-family: 'Titillium Web';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPDcZTIAOhVxoMyOr9n_E7ffHjDGIVzY4SY.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Titillium Web';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/s/titilliumweb/v15/NaPDcZTIAOhVxoMyOr9n_E7ffHjDGItzYw.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* cyrillic */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZC7-Am2FI_-VOxaLDvvq27p1O91Ko&skey=cf7b06949deecd8b&v=v19) format('woff2');
  unicode-range: U+0301, U+0400-045F, U+0490-0491, U+04B0-04B1, U+2116;
}

/* vietnamese */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZC7-Am2FI_-VOxaLDvvq27p1O21Ko&skey=cf7b06949deecd8b&v=v19) format('woff2');
  unicode-range: U+0102-0103, U+0110-0111, U+0128-0129, U+0168-0169, U+01A0-01A1, U+01AF-01B0, U+1EA0-1EF9, U+20AB;
}

/* latin-ext */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZC7-Am2FI_-VOxaLDvvq27p1O31Ko&skey=cf7b06949deecd8b&v=v19) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZC7-Am2FI_-VOxaLDvvq27p1O51Ko&skey=cf7b06949deecd8b&v=v19) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* cyrillic */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZB7-Am2FI_-VOxaLDvvq27p1sC8b_AsWc&skey=b4833ce002cd81b5&v=v19) format('woff2');
  unicode-range: U+0301, U+0400-045F, U+0490-0491, U+04B0-04B1, U+2116;
}

/* vietnamese */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZB7-Am2FI_-VOxaLDvvq27p1sC8b_LsWc&skey=b4833ce002cd81b5&v=v19) format('woff2');
  unicode-range: U+0102-0103, U+0110-0111, U+0128-0129, U+0168-0169, U+01A0-01A1, U+01AF-01B0, U+1EA0-1EF9, U+20AB;
}

/* latin-ext */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZB7-Am2FI_-VOxaLDvvq27p1sC8b_KsWc&skey=b4833ce002cd81b5&v=v19) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Proxima Nova';
  font-style: italic;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZB7-Am2FI_-VOxaLDvvq27p1sC8b_EsWc&skey=b4833ce002cd81b5&v=v19) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* cyrillic */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZc7-Am2FI_-VOxaLDvvq27pmO7&skey=96867d716c89840e&v=v19) format('woff2');
  unicode-range: U+0301, U+0400-045F, U+0490-0491, U+04B0-04B1, U+2116;
}

/* vietnamese */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZc7-Am2FI_-VOxaLDvvq27rWO7&skey=96867d716c89840e&v=v19) format('woff2');
  unicode-range: U+0102-0103, U+0110-0111, U+0128-0129, U+0168-0169, U+01A0-01A1, U+01AF-01B0, U+1EA0-1EF9, U+20AB;
}

/* latin-ext */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZc7-Am2FI_-VOxaLDvvq27rGO7&skey=96867d716c89840e&v=v19) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZc7-Am2FI_-VOxaLDvvq27omO7&skey=96867d716c89840e&v=v19) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

/* cyrillic */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZD7-Am2FI_-VOxaLDvvq2zGUau94_G&skey=470ea9b7edd795d7&v=v19) format('woff2');
  unicode-range: U+0301, U+0400-045F, U+0490-0491, U+04B0-04B1, U+2116;
}

/* vietnamese */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZD7-Am2FI_-VOxaLDvvq2zGUau_I_G&skey=470ea9b7edd795d7&v=v19) format('woff2');
  unicode-range: U+0102-0103, U+0110-0111, U+0128-0129, U+0168-0169, U+01A0-01A1, U+01AF-01B0, U+1EA0-1EF9, U+20AB;
}

/* latin-ext */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZD7-Am2FI_-VOxaLDvvq2zGUau_Y_G&skey=470ea9b7edd795d7&v=v19) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

/* latin */
@font-face {
  font-family: 'Proxima Nova';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/l/font?kit=FwZD7-Am2FI_-VOxaLDvvq2zGUau84_G&skey=470ea9b7edd795d7&v=v19) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
</style>
