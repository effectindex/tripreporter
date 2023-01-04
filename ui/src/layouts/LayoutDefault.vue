<template>
  <div class="LayoutDefault">
    <nav class="LayoutDefault__nav">
      <router-link to="/">Home</router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/login">Login</router-link>
    </nav>
    <main class="LayoutDefault__main">
      <slot/>
    </main>
    <footer class="LayoutDefault__footer">
      <div class="LayoutDefault__footer_text">
        Copyright &copy; 2022, Subjective Effect Documentation, 5HT2 and contributors.
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: "LayoutDefault"
}
</script>

<style scoped>
.LayoutDefault {
  font-family: "Titillium Web",-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.LayoutDefault__nav {
  font-family: "Proxima Nova",-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;
  background-color: #2e2e2e;
  padding: 30px;
  text-align: right;
}

.LayoutDefault__nav a {
  color: #ccc;
  text-decoration: none;
  text-transform: uppercase;
  font-weight: bold;
  margin-left: 1em;
}

.LayoutDefault__nav a.router-link-exact-active {
  color: #3d9991;
}

.LayoutDefault__nav > a:hover {
  color: #fff;
  text-shadow: 0 0 2px #96dbfc;
}

/* Used to pin footer to bottom */
.LayoutDefault__footer {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 60px;
}

/* Used to vertical center footer text */
.LayoutDefault__footer_text {
  text-align: center;
  position: relative;
  top: 50%;
  -ms-transform: translateY(-50%);
  -webkit-transform: translateY(-50%);
  transform: translateY(-50%);
}
</style>
