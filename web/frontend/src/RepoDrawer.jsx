import React, { useEffect } from "react";

import Drawer from "@material-ui/core/Drawer";
import Toolbar from "@material-ui/core/Toolbar";
import { CircularProgress, Grid, makeStyles } from "@material-ui/core";

import { useQuery, gql, useReactiveVar } from "@apollo/client";

import RepoInfoCard from "./RepoInfoCard";
import { repoNameVar } from "./ApolloClient";

const drawerWidth = 400;

const useStyle = makeStyles((theme) => ({
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  paper: {
    backgroundColor: theme.palette.background.default,
  },
  grid: {
    width: drawerWidth,
    padding: theme.spacing(2),
  },
}));

const GET_REPO_NAMES = gql`
  query GetRepoNames {
    repos {
      name
    }
  }
`;

export default function RepoDrawer() {
  const classes = useStyle();
  const { loading, error, data } = useQuery(GET_REPO_NAMES);
  const selectedRepoName = useReactiveVar(repoNameVar);

  useEffect(() => {
    if (repoNameVar() === "" && !loading && !error && data.repos.length > 0) {
      repoNameVar(data.repos[0].name);
    }
  });

  const renderRepoCards = () => {
    if (loading) {
      return (
        <Grid container item xs justify="center" alignItems="center">
          <CircularProgress />
        </Grid>
      );
    }
    if (error) {
      return <></>;
    }
    return data.repos.map((r) => {
      return (
        <Grid item xs key={r.name}>
          <RepoInfoCard
            repoName={r.name}
            selected={selectedRepoName === r.name}
          />
        </Grid>
      );
    });
  };

  return (
    <Drawer
      classes={{ paper: classes.paper }}
      className={classes.drawer}
      variant="permanent"
    >
      <Toolbar />
      <Grid container className={classes.grid} spacing={2} direction="column">
        {renderRepoCards()}
      </Grid>
    </Drawer>
  );
}
