import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardActionArea from "@material-ui/core/CardActionArea";

import PropTypes from "prop-types";

import { repoNameVar } from "./ApolloClient";

const useStyles = makeStyles((theme) => ({
  card: ({ selected }) => ({
    backgroundColor: selected
      ? theme.palette.primary.main
      : theme.palette.common.white,
    color: selected
      ? theme.palette.primary.contrastText
      : theme.palette.text.primary,
  }),
  title: ({ selected }) => ({
    fontSize: 14,
    color: selected
      ? theme.palette.primary.contrastText
      : theme.palette.text.secondary,
  }),
}));

export default function RepoInfoCard({ repoName, selected }) {
  const classes = useStyles({ selected });

  return (
    <Card className={classes.card}>
      <CardActionArea onClick={() => repoNameVar(repoName)}>
        <CardContent>
          <Typography className={classes.title} gutterBottom>
            Git Repo
          </Typography>
          <Typography variant="h5" component="h2">
            {repoName}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

RepoInfoCard.propTypes = {
  repoName: PropTypes.string.isRequired,
  selected: PropTypes.bool.isRequired,
};
