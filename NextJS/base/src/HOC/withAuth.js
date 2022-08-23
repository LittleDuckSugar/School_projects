import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";

const withAuth = (WrappedComponent) => {
	// eslint-disable-next-line react/display-name
	return (props) => {
		const Router = useRouter();
		const [authVerified, setAuthVerified] = useState(false);

		useEffect(() => {
			const token = localStorage.getItem("token");
			if (!token) {
				Router.push("/login");
			} else {
				setAuthVerified(true);
			}
		}, [Router]);
		if (authVerified) {
			return <WrappedComponent {...props} />;
		}
		else {
			return null;
		}
	};
};

export default withAuth;
