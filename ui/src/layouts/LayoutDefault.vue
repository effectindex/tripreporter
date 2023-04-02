<template>
  <div class="LayoutDefault">
    <div class="LayoutDefault__nav_box">
      <nav class="LayoutDefault__nav">
        <router-link class="LayoutDefault__item" to="/">Home</router-link>
        <router-link class="LayoutDefault__item" to="/about">About</router-link>
        <router-link v-if="!store.activeSession" class="LayoutDefault__item" to="/signup">Signup</router-link>
        <router-link v-if="!store.activeSession" class="LayoutDefault__item" to="/login">Login</router-link>
        <router-link v-if="store.activeSession" class="LayoutDefault__item" to="/create">Create Report</router-link>
      </nav>
    </div>
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
        <div class="LayoutDefault__footer_text" :set="year = getYear()">
          &copy; {{ year }}, Subjective Effect Documentation, 5HT2 and contributors.
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "LayoutDefault",
  methods: {
    getYear() {
      return new Date().getFullYear();
    }
  }
}
</script>

<script setup>
import {inject} from "vue";
import {useSessionStore} from '@/assets/lib/sessionstore'

const axios = inject('axios');
const store = useSessionStore();

store.updateSession(axios);
</script>

<style>
@import url(@/assets/css/default.css);
</style>

<style scoped>
@import url(@/assets/css/fonts.css);

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

/* make navbar float on top */
.LayoutDefault__nav_box {
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 1; /* fixes items rendering under navbar */
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

/* push main view down to avoid being hidden under navbar + pad bottom */
.LayoutDefault__main {
  padding: 90px 0;
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
  font-size: 0.75em;
}

/* override LayoutDefault__footer_text for desktop browsers */
@media only screen and (min-width: 680px) {
    .LayoutDefault__footer_text {
        position: relative;
        top: 50%;
        transform: translateY(-50%);
        font-size: revert;
    }
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
</style>
