<template>
  <div class="flex flex-row justify-center">
    <section class="max-w-lg mt-72">
      <div v-if="first_time" class="text-2xl">Create User</div>
      <div v-else class="text-2xl">Login</div>
      <o-field label="Username">
        <o-input type="text" @keydown.enter="submit" v-model="username" />
      </o-field>
      <o-field label="Password">
        <o-input type="password" @keydown.enter="submit" v-model="password" />
      </o-field>
      <o-field v-if="first_time" label="Confirm Password">
        <o-input
          type="password"
          @keydown.enter="submit"
          v-model="confirmPassword"
        />
      </o-field>
      <div class="flex flex-row justify-between mt-2">
        <div />
        <o-button variant="primary" v-if="first_time" @click="submit"
          >Register</o-button
        >
        <o-button variant="primary" v-else @click="submit">Login</o-button>
      </div>
    </section>
    <o-loading
      :full-page="true"
      v-model:active="loading"
      :can-cancel="false"
    ></o-loading>
  </div>
</template>

<script>
import APIUtil from "../util/APIUtil";
import AuthUtil from "../util/AuthUtil";

export default {
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
      loading: true,
      first_time: false,
    };
  },
  methods: {
    submit() {
      if (this.first_time) {
        this.submitRegister();
      } else {
        this.submitLogin();
      }
    },
    submitLogin() {
      this.loading = true;
      APIUtil.signIn(this.username, this.password)
        .then(() => {
          this.$router.push({ name: "home" });
          this.$store.commit("setLoggedIn", true);
        })
        .catch((msg) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Error logging in: ${msg}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
          this.$store.commit("setLoggedIn", false);
        })
        .finally(() => {
          this.loading = false;
        });
    },
    submitRegister() {
      if (this.confirmPassword != this.password) {
        this.$oruga.notification.open({
          duration: 5000,
          message: `Passwords do not match`,
          position: "bottom-right",
          variant: "danger",
          closable: false,
        });
        return;
      }
      this.loading = true;
      APIUtil.signUp(this.username, this.password)
        .then(() => {
          this.$router.push({ name: "home" });
        })
        .catch((msg) => {
          this.$oruga.notification.open({
            duration: 5000,
            message: `Error signing up: ${msg}`,
            position: "bottom-right",
            variant: "danger",
            closable: false,
          });
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
  mounted() {
      if (this.$store.getters.loggedIn) {
        this.$router.push({ name: "home" });
      } else {
        APIUtil.isFirstTime()
          .then(() => {
            this.first_time = true;
          })
          .catch(() => {
            this.first_time = false;
          })
          .finally(() => {
            this.loading = false;
          });
      }
  },
};
</script>
