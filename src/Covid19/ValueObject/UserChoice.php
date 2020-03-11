<?php

namespace Covid19\ValueObject;

class UserChoice
{
    const allowedTypes = ['regione', 'provincia'];

    protected $type;
    protected $value;

    public function __construct(string $type, string $value)
    {
        if (!in_array($type, self::allowedTypes)) {
            throw new \UnexpectedValueException(sprintf('The type "%s" is not a valid type', $type));
        }
        if (empty($value)) {
            throw new \UnexpectedValueException('The value cannot be empty');
        }

        $this->type  = $type;
        $this->value = $value;
    }

    public function type()
    {
        return $this->type;
    }

    public function value()
    {
        return $this->value;
    }
}
