import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import CheckIcon from "@material-ui/icons/Check";
import CloseIcon from "@material-ui/icons/Close";
import UpdateIcon from "@material-ui/icons/Update";
import TodayIcon from "@material-ui/icons/Today";
import ScheduleIcon from "@material-ui/icons/Schedule";

import { SourceCommit, SourceBranch } from "mdi-material-ui";

import { formatDistanceToNow, formatDuration } from "date-fns";

import { useQuery, gql } from "@apollo/client";

import PropTypes from "prop-types";

import { CodeBlock, monokai } from "react-code-blocks";

const useStyles = makeStyles((theme) => ({
  card: {
    display: "flex",
    width: "100%",
  },
  statusInd: {
    backgroundColor: ({ data }) => {
      if (data === undefined) return theme.palette.error.light;
      const {
        repos: [
          {
            branchesConnection: {
              edges: [
                {
                  node: {
                    commit: {
                      runsConnection: {
                        edges: [
                          {
                            node: { status },
                          },
                        ],
                      },
                    },
                  },
                },
              ],
            },
          },
        ],
      } = data;
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
    width: "5px",
    height: "100%",
  },
  cardContent: {
    width: "100%",
    display: "flex",
  },
  statusIcon: {
    marginTop: "4px",
    color: ({ data }) => {
      if (data === undefined) return theme.palette.error.light;
      const {
        repos: [
          {
            branchesConnection: {
              edges: [
                {
                  node: {
                    commit: {
                      runsConnection: {
                        edges: [
                          {
                            node: { status },
                          },
                        ],
                      },
                    },
                  },
                },
              ],
            },
          },
        ],
      } = data;
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
  },
  grid: {
    paddingLeft: theme.spacing(2),
  },
  titleLine: {
    color: ({ data }) => {
      if (data === undefined) return theme.palette.error.light;
      const {
        repos: [
          {
            branchesConnection: {
              edges: [
                {
                  node: {
                    commit: {
                      runsConnection: {
                        edges: [
                          {
                            node: { status },
                          },
                        ],
                      },
                    },
                  },
                },
              ],
            },
          },
        ],
      } = data;
      switch (status) {
        case "IN_PROGRESS":
          return theme.palette.warning.light;
        case "FAILED":
          return theme.palette.error.light;
        case "SUCCEED":
        default:
          return theme.palette.success.light;
      }
    },
  },
  codeBlock: {
    width: "100%",
  },
}));

const GET_REPO_CURRENT_COMMIT = gql`
  query GetRepoNames($repoName: String) {
    repos(name: $repoName) {
      name
      branchesConnection(first: 1) {
        edges {
          node {
            name
            commit {
              hash
              msg
              runsConnection(first: 1) {
                edges {
                  node {
                    num
                    startTimestamp
                    duration
                    status
                    log
                  }
                }
              }
            }
          }
        }
      }
    }
  }
`;

export default function CurrentRun({ repoName }) {
  const { loading, error, data } = useQuery(GET_REPO_CURRENT_COMMIT, {
    variables: { repoName },
  });
  const classes = useStyles({ data });

  // check if data is valid and get all needed fields
  if (loading || error) {
    return <></>;
  }

  const {
    repos: [
      {
        branchesConnection: {
          edges: [
            {
              node: {
                name: branchName,
                commit: {
                  hash: commitID,
                  msg: commitMsg,
                  runsConnection: {
                    edges: [
                      {
                        node: {
                          num: runID,
                          startTimestamp,
                          duration,
                          status,
                          log,
                        },
                      },
                    ],
                  },
                },
              },
            },
          ],
        },
      },
    ],
  } = data;

  return (
    <Grid item container xs direction="column" spacing={2}>
      <Grid item container xs>
        <Card className={classes.card}>
          <Paper elevation={0} className={classes.statusInd} />
          <CardContent className={classes.cardContent}>
            {((s) => {
              switch (s) {
                case "IN_PROGRESS":
                  return <UpdateIcon className={classes.statusIcon} />;
                case "FAILED":
                  return <CloseIcon className={classes.statusIcon} />;
                case "SUCCEED":
                default:
                  return <CheckIcon className={classes.statusIcon} />;
              }
            })(status)}
            <Grid container className={classes.grid} spacing={2}>
              <Grid
                item
                container
                xs={6}
                alignItems="center"
                className={classes.titleLine}
              >
                <Typography variant="h5">
                  <b>{branchName}</b>
                </Typography>
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  {commitMsg}
                </Typography>
              </Grid>
              <Grid
                item
                container
                xs={6}
                alignItems="center"
                className={classes.titleLine}
              >
                <SourceCommit />
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  <b>#{runID}</b>
                </Typography>
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  passed
                </Typography>
              </Grid>
              <Grid item container xs={6} alignItems="center">
                <SourceCommit />
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  Commit
                </Typography>
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  {commitID}
                </Typography>
              </Grid>
              <Grid item container xs={6} alignItems="center">
                <ScheduleIcon />
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  {formatDuration({
                    hours: Math.floor(duration / 3600),
                    minutes: Math.floor(duration / 60) % 60,
                    seconds: duration % 60,
                  })}
                </Typography>
              </Grid>
              <Grid item container xs={6} alignItems="center">
                <SourceBranch />
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  Branch
                </Typography>
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  {branchName}
                </Typography>
              </Grid>
              <Grid item container xs={6} alignItems="center">
                <TodayIcon />
                <Typography variant="h5">
                  <span>&nbsp;&nbsp;</span>
                  {formatDistanceToNow(new Date(startTimestamp) * 1000, {
                    addSuffix: true,
                  })}
                </Typography>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      </Grid>
      <Grid item container xs>
        <div className={classes.codeBlock}>
          <CodeBlock text={log} language="text" theme={monokai} />
        </div>
      </Grid>
    </Grid>
  );
}

CurrentRun.propTypes = {
  repoName: PropTypes.string.isRequired,
};
