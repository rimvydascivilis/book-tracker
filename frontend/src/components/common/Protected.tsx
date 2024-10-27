import React from "react";
import {Navigate} from "react-router-dom";
import PathConstants from "../../routes/PathConstants";
import {useAuth} from "../../context/AuthContext";

interface ProtectedRouteProps {
  children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({children}) => {
  const {isAuthenticated} = useAuth();

  if (!isAuthenticated) {
    return <Navigate to={PathConstants.LOGIN} />;
  }
  return <>{children}</>;
};

export default ProtectedRoute;