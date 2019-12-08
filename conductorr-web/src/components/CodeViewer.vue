<template>
    <div class="wrapper">
        <div class="code">{{ codeText }}</div>
        <div
            v-clipboard:copy="codeText"
            v-clipboard:success="copySuccess"
            v-clipboard:error="copyFail"
            class="copy"
        >
            <b-icon icon="copy" size="is-medium"> </b-icon>
        </div>
    </div>
</template>

<script>
export default {
    name: "code-viewer",
    components: {},
    props: ["codeText"],
    methods: {
        copySuccess(){
            this.$buefy.toast.open('Copied to clipboard')
        },
        copyFail(){
            this.$buefy.toast.open('Error copying')
        }
    }
};
</script>

<style scoped>
.code {
    font-family: monospace;
    width: 100%;
    background: rgb(240, 240, 240);
    border-radius: 5px;
    height: 400px;
    overflow-y: scroll;
    position: relative;
}
.wrapper {
    position: relative;
}
.copy {
    cursor: pointer;
    position: absolute;
    top: 5px;
    right: 18px;
}
</style>
