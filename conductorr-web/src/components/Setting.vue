<template>
    <div class="tile is-parent" style="margin-top:20px;">
        <div class="tile is-4">
            <b-field :label="setting.label" v-if="setting.datatype == 'string'">
                <b-input :placeholder="setting.placeholder" v-model="setting.string_value"></b-input>
            </b-field>
            <div class="field" v-if="setting.datatype == 'bool'">
                <b-switch size="is-medium" :value="setting.bool_value" type="is-success" controls-position="compact">
                    {{setting.label}}
                </b-switch>
            </div>
            <b-field :label="setting.label" v-if="setting.datatype == 'password'">
                <b-input v-bind:value="setting.string_value" type="password" maxlength="30"></b-input>
            </b-field>
            <b-field :label="setting.label" v-if="setting.datatype == 'int'">
                <b-numberinput v-model="setting.int_value"></b-numberinput>
            </b-field>
        </div>
        <div class="tile is-1"></div>
        <div class="tile">
            <b-message type="is-info">
                {{setting.description}}
            </b-message>
        </div>
    </div>
</template>

<script>
export default {
    name: 'setting',
    props: [
        'setting'
    ]
}
</script>