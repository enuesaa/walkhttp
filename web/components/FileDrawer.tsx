import { Drawer, DrawerBody, DrawerFooter, DrawerOverlay, DrawerContent, DrawerCloseButton, useDisclosure, Button} from '@chakra-ui/react'
import { useRef } from 'react'

export const FileDrawer = () => {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const btnRef = useRef<HTMLButtonElement>(null)

  return (
    <>
      <Button ref={btnRef} colorScheme='purple' onClick={onOpen}>
        Change Files
      </Button>
      <Drawer
        isOpen={isOpen}
        placement='right'
        onClose={onClose}
        finalFocusRef={btnRef}
      >
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerBody>
            a
          </DrawerBody>
          <DrawerFooter>
            <Button colorScheme='purple'>Open</Button>
          </DrawerFooter>
        </DrawerContent>
      </Drawer>
    </>
  )
}