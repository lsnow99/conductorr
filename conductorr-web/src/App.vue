<template>
    <!--  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link> |
      <router-link to="/configuration">Settings</router-link> |
      <router-link to="/overview">Overview</router-link>
    </div>
  </div> -->
    <div id="app">
        <b-navbar>
            <template slot="brand">
                <b-navbar-item tag="router-link" :to="{ path: '/' }">
                    <img width="150px" src="@/assets/default-monochrome.svg"/>
                </b-navbar-item>
            </template>
            <template slot="start">
                <b-navbar-item tag="router-link" to="/">
                    Home
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/statistics">
                    Statistics
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/configuration">
                    Configuration
                </b-navbar-item>
            </template>

            <template slot="end">
                <b-navbar-item tag="div">
                    <div class="buttons" v-if="loggedIn()">
                        <a @click="logout"  class="button is-light">
                            Log out
                        </a>
                    </div>
                </b-navbar-item>
            </template>
        </b-navbar>
        <div id="mainView">
            <router-view />
        </div>
    </div>
</template>

<style>
#mainView {
    margin-top: 20px;
}

#app {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
}

#nav {
    padding: 30px;
}

#nav a {
    font-weight: bold;
    color: #2c3e50;
}

#nav a.router-link-exact-active {
    color: #42b983;
}
</style>

<script>
export default {
    components: {}
};
</script>

<style lang="scss">
// Import Bulma's core
@import "~bulma/sass/utilities/_all";

// Set your colors
$primary: #eca02d;
$primary-invert: findColorInvert($primary);
$twitter: #4099ff;
$twitter-invert: findColorInvert($twitter);
$health: $red;
$health-invert: findColorInvert($health);

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
    "white": (
        $white,
        $black
    ),
    "black": (
        $black,
        $white
    ),
    "light": (
        $light,
        $light-invert
    ),
    "dark": (
        $dark,
        $dark-invert
    ),
    "primary": (
        $primary,
        $primary-invert
    ),
    "info": (
        $info,
        $info-invert
    ),
    "success": (
        $success,
        $success-invert
    ),
    "warning": (
        $warning,
        $warning-invert
    ),
    "danger": (
        $danger,
        $danger-invert
    ),
    "twitter": (
        $twitter,
        $twitter-invert
    ),
    "health": (
        $health,
        $health-invert
    )
);

// Links
$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

// Import Bulma and Buefy styles
@import "~bulma";
@import "~buefy/src/scss/buefy";
</style>
