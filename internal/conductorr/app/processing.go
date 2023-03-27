package app

import (
  "io"
  "os"
	"github.com/rs/zerolog/log"
)

type ProcessingJob struct {
  fromPath string
  toPath string
  copy bool
  callback func(err error)
  downloadId int
}

type ProcessingDistributor struct {
  jobChan chan ProcessingJob
}

var Processor *ProcessingDistributor

func ProcessJob(job ProcessingJob) {
  var err error
  defer func() {
    job.callback(err)
  }()
  srcFile, err := os.OpenFile(job.fromPath, os.O_RDONLY, os.ModePerm)
  if err != nil {
    log.Error().
      Err(err).
      Str("video_path", job.fromPath).
      Int("download_id", job.downloadId).
      Msg("could not open source file")

    return
  }
  if job.copy {
    destFile, err := os.OpenFile(job.toPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
    if err != nil {
      log.Error().
        Err(err).
        Str("dest_filepath", job.toPath).
        Int("download_id", job.downloadId).
        Msg("could not open output file")
      return
    }
    n, err := io.Copy(destFile, srcFile)
    if err != nil {
      log.Error().
        Err(err).
        Int("download_id", job.downloadId).
        Msgf("got error after copying %d bytes", n)
      return
    }
  } else {
    err = os.Rename(job.fromPath, job.toPath)
    if err != nil {
      log.Error().Err(err).Int("download_id", job.downloadId).Msg("error moving file")
    }
  }
}

func (pd *ProcessingDistributor) spawnWorker() {
  for {
    select {
      case job := <-pd.jobChan:
        // Process job
        ProcessJob(job)
    }
  }
}

func (pd *ProcessingDistributor) AddJob(job ProcessingJob) {
  pd.jobChan <- job
}

func NewProcessor(numWorkers int) *ProcessingDistributor {
  pd := ProcessingDistributor{
    jobChan: make(chan ProcessingJob, 500),
  }
  for i := 0; i < numWorkers; i++ {
    go pd.spawnWorker()
  }
  return &pd
}

func init() {
  Processor = NewProcessor(3)
}
