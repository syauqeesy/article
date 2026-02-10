<script lang="ts">
  import { onMount } from "svelte";
  import { Editor } from "@tiptap/core";
  import StarterKit from "@tiptap/starter-kit";
  import Image from "@tiptap/extension-image";

  let editor: Editor | null = null;
  let editorElement: HTMLDivElement | null = null;
  let fileInputElement: HTMLInputElement | null = null;

  onMount(() => {
    editor = new Editor({
      element: editorElement,
      extensions: [StarterKit, Image],
    });

    return () => {
      if (editor) editor.destroy();
    };
  });

  const save = () => {
    if (!editor) return console.error("editor not instantiated");

    console.log(editor.getJSON());
  };

  const handleFileInputChange = (event: Event) => {
    if (!editor) return console.error("editor not instantiated");

    const target = event.target as HTMLInputElement;

    const file = target.files?.[0];
    if (file)
      editor
        .chain()
        .focus()
        .setImage({
          src: URL.createObjectURL(file),
        })
        .run();
  };

  const triggerFileInput = () => {
    if (!fileInputElement) {
      return console.error("file input not instantiated");
    }

    (fileInputElement as unknown as HTMLInputElement).click();
  };
</script>

<div class="editor">
  <div class="editor-toolbar">
    <input
      type="file"
      name="image"
      id="image"
      accept="image/*"
      bind:this={fileInputElement}
      on:change={handleFileInputChange}
    />
    <button on:click={triggerFileInput}>image</button>
  </div>
  <div class="editor-content" bind:this={editorElement}></div>
  <button on:click={save}>save</button>
</div>

<style>
  div.editor {
    padding: 1rem;
  }

  div.editor div {
    border: 1px solid #ddd;
    padding: 1rem;
    margin-bottom: 0.5rem;
    border-radius: 0.5rem;
  }

  div.editor div.editor-content {
    min-height: 200px;
  }

  div.editor div.editor-content :global(.ProseMirror) {
    outline: none;
  }

  div.editor div.editor-content :global(.ProseMirror) {
    min-height: inherit;
    height: inherit;
    box-sizing: border-box;
  }

  div.editor div.editor-content :global(.ProseMirror img) {
    max-width: 100%;
    height: auto;
    display: block;
    margin: 0.5rem auto;
  }

  div.editor div input[name="image"] {
    display: none;
  }
</style>
