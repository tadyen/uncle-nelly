package unclenelly

import (
    "fmt"
    "errors"
)

type JobName string

const (
    CookingSim  JobName = "CookingSim"
    ReverseCook JobName = "ReverseCook"
    Optimise    JobName = "Optimise"
)

func (j JobName) String() string {
    return string(j)
}

func (j JobName) Validate() error {
    switch j {
    case CookingSim,
        ReverseCook,
        Optimise:
        return nil
    default:
        return fmt.Errorf("Wrong operation name: %s", j)
    }
}

// config
type Job struct{
    // CookingSim | ReverseCook | Optimise
    Name    JobName
    Product *Product
}

func NewJob(jobName string) (*Job, error) {
    job := DefaultJob()
    if jobName == "" {
        return job, nil
    }
    err := JobName(jobName).Validate()
    if err != nil {
        return nil, err
    }
    job.Name = JobName(jobName)
    return job, nil
}

func DefaultJob() *Job {
    newProduct, _ := NewProduct(BlankBaseIngredient)
    return &Job{
        Name:       CookingSim,
        Product:    newProduct,
    }
}


// Validate if job is valid, return nil if valid
func(j *Job) Validate() error {
    var errs []error
    var err error
    err = j.Name.Validate()
    if err != nil {
        errs = append(errs, err)
    }
    return errors.Join(errs...)
}

