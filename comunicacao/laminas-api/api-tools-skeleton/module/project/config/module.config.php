<?php
return [
    'router' => [
        'routes' => [
            'project.rest.users' => [
                'type' => 'Segment',
                'options' => [
                    'route' => '/users[/:users_id]',
                    'defaults' => [
                        'controller' => 'project\\V1\\Rest\\Users\\Controller',
                    ],
                ],
            ],
        ],
    ],
    'api-tools-versioning' => [
        'uri' => [
            0 => 'project.rest.users',
        ],
    ],
    'api-tools-rest' => [
        'project\\V1\\Rest\\Users\\Controller' => [
            'listener' => 'project\\V1\\Rest\\Users\\UsersResource',
            'route_name' => 'project.rest.users',
            'route_identifier_name' => 'users_id',
            'collection_name' => 'users',
            'entity_http_methods' => [
                0 => 'GET',
                1 => 'PATCH',
                2 => 'PUT',
                3 => 'DELETE',
            ],
            'collection_http_methods' => [
                0 => 'GET',
                1 => 'POST',
            ],
            'collection_query_whitelist' => [],
            'page_size' => 25,
            'page_size_param' => null,
            'entity_class' => \project\V1\Rest\Users\UsersEntity::class,
            'collection_class' => \project\V1\Rest\Users\UsersCollection::class,
            'service_name' => 'users',
        ],
    ],
    'api-tools-content-negotiation' => [
        'controllers' => [
            'project\\V1\\Rest\\Users\\Controller' => 'HalJson',
        ],
        'accept_whitelist' => [
            'project\\V1\\Rest\\Users\\Controller' => [
                0 => 'application/vnd.project.v1+json',
                1 => 'application/hal+json',
                2 => 'application/json',
            ],
        ],
        'content_type_whitelist' => [
            'project\\V1\\Rest\\Users\\Controller' => [
                0 => 'application/vnd.project.v1+json',
                1 => 'application/json',
            ],
        ],
    ],
    'api-tools-hal' => [
        'metadata_map' => [
            \project\V1\Rest\Users\UsersEntity::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'project.rest.users',
                'route_identifier_name' => 'users_id',
                'hydrator' => \Laminas\Hydrator\ArraySerializableHydrator::class,
            ],
            \project\V1\Rest\Users\UsersCollection::class => [
                'entity_identifier_name' => 'id',
                'route_name' => 'project.rest.users',
                'route_identifier_name' => 'users_id',
                'is_collection' => true,
            ],
        ],
    ],
    'api-tools' => [
        'db-connected' => [
            'project\\V1\\Rest\\Users\\UsersResource' => [
                'adapter_name' => 'dummy',
                'table_name' => 'users',
                'hydrator_name' => \Laminas\Hydrator\ArraySerializableHydrator::class,
                'controller_service_name' => 'project\\V1\\Rest\\Users\\Controller',
                'entity_identifier_name' => 'id',
            ],
        ],
    ],
    'api-tools-content-validation' => [
        'project\\V1\\Rest\\Users\\Controller' => [
            'input_filter' => 'project\\V1\\Rest\\Users\\Validator',
        ],
    ],
    'input_filter_specs' => [
        'project\\V1\\Rest\\Users\\Validator' => [
            0 => [
                'name' => 'id',
                'required' => false,
                'filters' => [
                    0 => [
                        'name' => \Laminas\Filter\StripTags::class,
                    ],
                    1 => [
                        'name' => \Laminas\Filter\Digits::class,
                    ],
                ],
                'validators' => [],
            ],
            1 => [
                'name' => 'name',
                'required' => false,
                'filters' => [],
                'validators' => [],
            ],
            2 => [
                'name' => 'email',
                'required' => true,
                'filters' => [],
                'validators' => [],
                'field_type' => 'email',
            ],
        ],
    ],
];
