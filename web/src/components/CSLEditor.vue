<template>
  <prism-editor
    class="my-editor"
    style="height: 100%"
    v-model="computedValue"
    :highlight="highlighter"
    :readonly="readonly"
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
  font-size: 14px;
  line-height: 1.5;
  padding: 5px;
}

/* optional class for removing the outline */
.prism-editor__textarea:focus {
  outline: none;
}

/* attempt to remove word-wrap */
.prism-editor__textarea {
  width: 999999px !important;
}

.prism-editor-wrapper .prism-editor__editor,
.prism-editor-wrapper .prism-editor__textarea {
  white-space: pre !important;
}

.prism-editor__container {
  overflow-x: scroll !important;
}
</style>

<script>
import { highlight, languages } from "prismjs/components/prism-core";
import "prismjs/components/prism-clike";
import "prismjs/components/prism-lisp";
import "prismjs/themes/prism-tomorrow.css"; // import syntax highlighting styles

export default {
  name: 'csl-editor',
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
  emits: ['update:modelValue'],
  methods: {
    highlighter(code) {
      return highlight(code, languages.lisp, "lisp");
    },
  },
  computed: {
    computedValue: {
      get() {
        return this.modelValue
      },
      set(newVal) {
        this.$emit('update:modelValue', newVal)
      }
    }
  }
};
</script>
