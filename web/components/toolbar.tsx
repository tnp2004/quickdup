import { Editor } from "@tiptap/react";

export default function Toolbar({ editor }: {editor: Editor | null}) {
  const isActive = (className: string): string => {
    return editor?.isActive(className) ? "font-bold" : ""
  }
    return (
      <div className="flex gap-2 mb-1 py-1">
        <button className={`${isActive("bold")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleBold().run()}>bold</button>
        <button className={`${isActive("heading")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleHeading({level: 3}).run()}>heading</button>
        <button className={`${isActive("italic")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleItalic().run()}>italic</button>
        <button className={`${isActive("list")} toolBtn`}
          onClick={() => editor?.chain().focus().toggleBulletList().run()}>list</button>
      </div>
    )
}