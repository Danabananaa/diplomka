import {
  createBrowserRouter, 
  createRoutesFromElements,
  redirect,
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
import { homeLoader } from './api/loaders';

const router = createBrowserRouter(
  createRoutesFromElements(
<Route>
  <Route element={<RootLayout />} id="root" errorElement={<ErrorPage/>}>
      <Route path="/" >
          <Route index 
            loader={homeLoader}
          />
          <Route path="statistics" element={<HomePage/>}/>              
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