import {
  createBrowserRouter, 
  createRoutesFromElements,
  Route, 
  RouterProvider,
} from 'react-router-dom'  
// pages
import SignIn from './routes/Sign-In/SignIn'
import SignUp from './routes/Sign-Up/SignUp';
import HomePage from './routes/Home-Page/HomePage'
// layouts
import RootLayout from './layouts/RootLayout/RootLayout';
// errorElement
import { ErrorPage } from './routes/Error-Page/ErrorPage';
import { PageNotFound } from './routes/Page-Not-Found/PageNotFound';

const router = createBrowserRouter(
  createRoutesFromElements(
<Route>
  <Route element={<RootLayout />} id="root" errorElement={<ErrorPage/>}>
      <Route path="/" >
          <Route index 
            element={<HomePage />}
          />              
      </Route>
      <Route path="*" element={<PageNotFound/>} />
  </Route>
      <Route path="signin" element={<SignIn/>} />
      <Route path="signup" element={<SignUp/>}  />
</Route>
  )
)

const App = () => {
  return (
    <RouterProvider router={router} />
  );
}

export default App;