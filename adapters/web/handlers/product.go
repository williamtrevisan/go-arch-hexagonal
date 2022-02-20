package handlers

import (
    "encoding/json"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/williamtrevisan/go-arch-hexagonal/adapters/dto"
    "github.com/williamtrevisan/go-arch-hexagonal/app"
    "net/http"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
    r.Handle("/product", n.With(
        negroni.Wrap(createProduct(service)),
    ))

    r.Handle("/product/{id}", n.With(
        negroni.Wrap(getProduct(service)),
    ))

    r.Handle("/product/{id}/disable", n.With(
        negroni.Wrap(disableProduct(service)),
    ))

    r.Handle("/product/{id}/enable", n.With(
        negroni.Wrap(enableProduct(service)),
    ))
}

func createProduct(service app.ProductServiceInterface) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        var productDto dto.Product

        err := json.NewDecoder(r.Body).Decode(&productDto)
        if err != nil {
            returnInternalServerError(w)

            return
        }

        product, err := service.Create(productDto.Name, productDto.Price)
        if err != nil {
            returnInternalServerError(w)

            return
        }

        err = json.NewEncoder(w).Encode(product)
        if err != nil {
            returnInternalServerError(w)

            return
        }
    })
}

func getProduct(service app.ProductServiceInterface) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        vars := mux.Vars(r)

        product, err := service.Get(vars["id"])
        if err != nil {
            w.WriteHeader(http.StatusNotFound)
            w.Write(jsonError(err.Error()))

            return
        }

        err = json.NewEncoder(w).Encode(product)
        if err != nil {
            returnInternalServerError(w)

            return
        }
    })
}

func enableProduct(service app.ProductServiceInterface) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        vars := mux.Vars(r)

        product, err := service.Get(vars["id"])
        if err != nil {
            w.WriteHeader(http.StatusNotFound)
            w.Write(jsonError(err.Error()))

            return
        }

        result, err := service.Enable(product)
        if err != nil {
            returnInternalServerError(w)

            return
        }

        err = json.NewEncoder(w).Encode(result)
        if err != nil {
            returnInternalServerError(w)

            return
        }
    })
}

func disableProduct(service app.ProductServiceInterface) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        vars := mux.Vars(r)

        product, err := service.Get(vars["id"])
        if err != nil {
            w.WriteHeader(http.StatusNotFound)
            w.Write(jsonError(err.Error()))

            return
        }

        result, err := service.Disable(product)
        if err != nil {
            returnInternalServerError(w)

            return
        }

        err = json.NewEncoder(w).Encode(result)
        if err != nil {
            returnInternalServerError(w)

            return
        }
    })
}

func returnInternalServerError(w http.ResponseWriter) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(jsonError(err.Error()))
}
