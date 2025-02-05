import { Container, Stack } from "@chakra-ui/react"
import { FcTodoList } from "react-icons/fc"

export default function App() {
  return (
    <Stack h="100vh">
     <Navbar/>
     <Container>
      <TodoForm/>
      <FcTodoList/>
     </Container>
    </Stack>
  )
}
