"use client";

import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";

export default function Editor() {
  const editor = useEditor({
    extensions: [StarterKit],
    content: "<p>Hello world</p>",
    editorProps: {
        attributes: {
            class: "min-h-[150px] p-1 focus:outline-none"
        }
    }
  });

  return (
    <div>
      <div className="w-1/2 mx-auto flex flex-col justify-center">
        <div className="border-2 rounded mb-1 py-1 px-3">
          <button
            className="rounded"
            onClick={() => editor?.chain().focus().toggleBold().run()}>bold</button>
        </div>
          <EditorContent
            className="w-full mx-auto border-2 rounded"
            editor={editor}
          />
      </div>
      <button onClick={() => console.log(editor?.getHTML())}>get html</button>
    </div>
  );
}
