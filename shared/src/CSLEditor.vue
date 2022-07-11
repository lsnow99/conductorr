<template>
  <PrismEditor
    ref="editor"
    aria-label="Edit CSL"
    label="edit csl"
    class="h-full my-editor"
    v-model="computedValue"
    @keydown.enter.ctrl="enterPressed"
    @click="focusEditor"
    :highlight="highlighter"
    :readonly="readonly"
    :aria-disabled="readonly"
    line-numbers
  ></PrismEditor>
</template>

<style scoped>
/* required class */
.my-editor {
  /* we dont use `language-` classes anymore so thats why we need to add background and text color manually */
  background: #2d2d2d;
  color: #ccc;

  /* you must provide font-family font-size line-height. Example: */
  font-family: Fira code, Fira Mono, Consolas, Menlo, Courier, monospace;
  font-size: 16px;
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

<script lang="ts">
export default {
  name: "csl-editor",
};
</script>

<script setup lang="ts">
import { PrismEditor } from "vue-prism-editor";
import { highlight, languages } from "prismjs/components/prism-core";
import "vue-prism-editor/dist/prismeditor.min.css";
import "prismjs/components/prism-clike";
import "prismjs/components/prism-lisp";
import "prismjs/themes/prism-tomorrow.css"; // import syntax highlighting styles
import { Component, ref } from "vue";
import { useComputedValue } from "@/util"

const props = defineProps<{
  modelValue: string;
  readonly?: boolean;
}>();

const emit = defineEmits<{
  (e: "update:modelValue", newVal: string): void;
  (e: "submit"): void;
}>();

const highlighter = (code: string) => highlight(code, languages.lisp, "lisp");

const enterPressed = ($event: Event) => {
  $event?.stopImmediatePropagation();
  emit("submit");
};

const editor = ref<Element>();
const focusEditor = () => {
  Array.from(editor.value?.getElementsByTagName("textarea")).forEach(
    (elem) => elem.focus()
  );
};

const computedValue = useComputedValue<string>(props, emit)
</script>
