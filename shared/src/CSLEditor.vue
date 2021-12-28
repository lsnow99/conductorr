<template>
  <prism-editor
    aria-label="Edit CSL"
    label="edit csl"
    class="my-editor h-full"
    v-model="computedValue"
    @keydown.enter.ctrl="enterPressed"
    @click="focusEditor"
    ref="editor"
    :highlight="highlighter"
    :readonly="readonly"
    :aria-disabled="readonly"
    line-numbers
  ></prism-editor>
</template>

<style scoped>
/* required class */
.my-editor {
  /* we dont use `language-` classes anymore so thats why we need to add background and text color manually */
  background: #2d2d2d;
  color: #ccc;

  /* you must provide font-family font-size line-height. Example: */
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 20px;
  line-height: 1.5;
  padding: 5px;
  cursor: text;
}

/* optional class for removing the outline */
/* :deep(.prism-editor__textarea:focus) {
  outline: none;
} */

/* attempt to remove word-wrap */
/* :deep(.prism-editor__textarea) {
  width: 5000px !important;
}
:deep(.prism-editor__editor) {
  white-space: pre !important;
}
:deep(.prism-editor__container) {
  overflow-x: scroll !important;
} */
</style>

<script>
import { PrismEditor } from "vue-prism-editor";
import { highlight, languages } from "prismjs/components/prism-core";
import "vue-prism-editor/dist/prismeditor.min.css";
import "prismjs/components/prism-clike";
import "prismjs/components/prism-lisp";
import "prismjs/themes/prism-tomorrow.css"; // import syntax highlighting styles

export default {
  name: "csl-editor",
  props: {
    modelValue: {
      type: String,
      default: function () {
        return "";
      },
    },
    readonly: {
      type: Boolean,
      default: function () {
        return false;
      },
    },
  },
  components: {
    PrismEditor,
  },
  emits: ["update:modelValue", "submit"],
  methods: {
    highlighter(code) {
      return highlight(code, languages.lisp, "lisp");
    },
    enterPressed(event) {
      event.stopImmediatePropagation();
      this.$emit("submit");
    },
    focusEditor() {
      const editors = Array.from(this.$el.getElementsByTagName("textarea"));
      if (editors.length < 1) {
        return;
      }
      editors[0].focus();
    },
  },
  computed: {
    computedValue: {
      get() {
        return this.modelValue;
      },
      set(newVal) {
        this.$emit("update:modelValue", newVal);
      },
    },
  },
};
</script>
