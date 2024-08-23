import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react'
import { Editor } from '@tiptap/react'
import { isActive } from '@/utils'

export default function HeadingDropdown({editor}: {editor: Editor | null}) {
  return (
    <Menu as="div" className="relative inline-block text-left">
      <div>
        <MenuButton className={`${isActive(editor,"heading")} inline-flex w-full justify-center gap-x-1.5 toolBtn`}>
          heading
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="lucide lucide-chevron-down w-5"><path d="m6 9 6 6 6-6"/></svg>
        </MenuButton>
      </div>

      <MenuItems
        transition
        className="absolute right-0 z-10 mt-2 w-fit origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 transition focus:outline-none data-[closed]:scale-95 data-[closed]:transform data-[closed]:opacity-0 data-[enter]:duration-100 data-[leave]:duration-75 data-[enter]:ease-out data-[leave]:ease-in"
      >
        <div className="py-1">
        
           <MenuItem>
            <button className={`dropdownElement`}
          onClick={() => editor?.chain().focus().toggleHeading({level: 1}).run()}>heading 1</button>
         </MenuItem>
           <MenuItem>
            <button className={`dropdownElement`}
          onClick={() => editor?.chain().focus().toggleHeading({level: 2}).run()}>heading 2</button>
         </MenuItem>
           <MenuItem>
            <button className={`dropdownElement`}
          onClick={() => editor?.chain().focus().toggleHeading({level: 3}).run()}>heading 3</button>
         </MenuItem>
        </div>
      </MenuItems>
    </Menu>
  )
}
