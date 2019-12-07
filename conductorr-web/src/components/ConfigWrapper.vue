<template>
    <div class="config-bg">
        <div class="controls">
            <div class="field">
                <b-switch :value="true"
                type="is-warning">
                    Show Advanced
                </b-switch>
            </div>
        </div>
        
        <slot></slot>
    </div>
</template>

<style scoped>
.config-bg{
    -webkit-box-shadow: 6px 6px 19px 0px rgba(168,168,168,1);
    -moz-box-shadow: 6px 6px 19px 0px rgba(168,168,168,1);
    box-shadow: 6px 6px 19px 0px rgba(168,168,168,1);

    background-color: rgb(250, 250, 250);
    border-radius: 7px;

    padding: 30px;
}
.controls{
    display: none;
}
.setting-header {
    font-size:25px;
}
</style>

<style>
.setting-header {
    font-size:30px;
    line-height: 50px;
}
.setting-logo {
    width: 50px;
    vertical-align: middle;
}
</style>

<script>
export default {
    name: 'ConfigWrapper'
}
</script>